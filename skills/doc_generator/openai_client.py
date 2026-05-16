"""Small OpenAI client wrapper for API DocAgent.

This module keeps OpenAI calls in one place so hackathon skills can reuse the
same environment loading, retries, logging, and response cleanup.
"""

from __future__ import annotations

import json
import logging
import os
import time
from pathlib import Path
from typing import Any

try:
    from dotenv import load_dotenv
except ImportError:  # pragma: no cover - local fallback for minimal setups
    load_dotenv = None  # type: ignore[assignment]

try:
    from openai import APIConnectionError, APIError, APITimeoutError, OpenAI, RateLimitError
except ImportError:  # pragma: no cover - lets the module import before deps install
    class _MissingOpenAIError(Exception):
        """Placeholder so config errors do not get retried as API failures."""

    APIConnectionError = APIError = APITimeoutError = RateLimitError = _MissingOpenAIError
    OpenAI = None  # type: ignore[assignment]


DEFAULT_MODEL = "gpt-4o-mini"
DEFAULT_MAX_RETRIES = 2
DEFAULT_TEMPERATURE = 0.2

logger = logging.getLogger(__name__)
logging.basicConfig(
    level=os.getenv("API_DOCAGENT_LOG_LEVEL", "INFO"),
    format="%(levelname)s:%(name)s:%(message)s",
)

_client: Any | None = None
_env_loaded = False


def _repo_root() -> Path:
    """Return the API DocAgent repository root from this skill file."""
    return Path(__file__).resolve().parents[2]


def _load_env_file_manually(env_path: Path) -> None:
    """Tiny .env fallback when python-dotenv is not installed."""
    if not env_path.exists():
        return

    for raw_line in env_path.read_text(encoding="utf-8").splitlines():
        line = raw_line.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue

        key, value = line.split("=", 1)
        clean_value = value.strip().strip('"').strip("'")
        os.environ.setdefault(key.strip(), clean_value)


def _load_environment() -> None:
    """Load OPENAI_API_KEY and optional settings from the repo .env file."""
    global _env_loaded

    if _env_loaded:
        return

    env_path = _repo_root() / ".env"
    if load_dotenv is not None:
        load_dotenv(env_path)
        load_dotenv()
    else:
        _load_env_file_manually(env_path)

    _env_loaded = True


def _get_max_retries() -> int:
    """Read retry count from env with a safe fallback."""
    try:
        return int(os.getenv("OPENAI_MAX_RETRIES", DEFAULT_MAX_RETRIES))
    except ValueError:
        logger.warning("Invalid OPENAI_MAX_RETRIES value; using default")
        return DEFAULT_MAX_RETRIES


def _get_client() -> Any:
    """Create and cache the OpenAI SDK client."""
    global _client

    if _client is not None:
        return _client

    _load_environment()

    if OpenAI is None:
        raise RuntimeError("OpenAI SDK is not installed. Run: pip install openai")

    api_key = os.getenv("OPENAI_API_KEY")
    if not api_key:
        raise RuntimeError("OPENAI_API_KEY is missing. Add it to .env or your shell environment.")

    _client = OpenAI(api_key=api_key)
    return _client


def _chat_completion(
    prompt: str,
    *,
    system_prompt: str,
    response_format: dict[str, str] | None = None,
) -> str:
    """Call OpenAI Chat Completions with simple retry handling."""
    max_retries = _get_max_retries()
    model = os.getenv("OPENAI_MODEL", DEFAULT_MODEL)
    retryable_errors = (APIConnectionError, APITimeoutError, RateLimitError, APIError)

    request_kwargs: dict[str, Any] = {
        "model": model,
        "temperature": DEFAULT_TEMPERATURE,
        "messages": [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": prompt},
        ],
    }
    if response_format is not None:
        request_kwargs["response_format"] = response_format

    for attempt in range(1, max_retries + 2):
        try:
            response = _get_client().chat.completions.create(**request_kwargs)
            content = response.choices[0].message.content or ""
            logger.info("Generated OpenAI response with model=%s", model)
            return content.strip()
        except retryable_errors as exc:
            if attempt > max_retries:
                logger.error("OpenAI request failed after %s attempts: %s", attempt, exc)
                raise

            sleep_seconds = min(0.5 * attempt, 2.0)
            logger.warning(
                "OpenAI request failed on attempt %s/%s; retrying in %.1fs: %s",
                attempt,
                max_retries + 1,
                sleep_seconds,
                exc,
            )
            time.sleep(sleep_seconds)

    return ""


def generate_text(prompt: str) -> str:
    """Generate a plain text response for documentation content."""
    try:
        return _chat_completion(
            prompt,
            system_prompt=(
                "You are API DocAgent. Generate concise, accurate API documentation "
                "from the provided Go service metadata."
            ),
        )
    except Exception as exc:  # pragma: no cover - defensive for hackathon CLI runs
        logger.error("Text generation failed: %s", exc)
        return ""


def generate_json(prompt: str) -> dict[str, Any]:
    """Generate and parse a structured JSON object response."""
    raw_content = ""

    try:
        raw_content = _chat_completion(
            prompt,
            system_prompt=(
                "You are API DocAgent. Return only one valid JSON object with no "
                "markdown fences, no prose, and no comments."
            ),
            response_format={"type": "json_object"},
        )
        parsed = json.loads(raw_content)

        if isinstance(parsed, dict):
            return parsed

        return {"data": parsed}
    except json.JSONDecodeError as exc:
        logger.error("Model returned invalid JSON: %s", exc)
        return {"error": "invalid_json", "raw": raw_content}
    except Exception as exc:  # pragma: no cover - defensive for hackathon CLI runs
        logger.error("JSON generation failed: %s", exc)
        return {"error": str(exc)}
