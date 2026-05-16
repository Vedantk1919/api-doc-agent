# API DocAgent Go Demo Service

Small isolated Go service for hackathon demos of API route discovery, request
and response schema extraction, documentation generation, validation, and drift
detection.

It mirrors the source repository shape at a much smaller scale:

- `nsdConfig/router.go` defines Gin routes.
- `app/controllers` binds requests and delegates to services.
- `app/services` performs validation and orchestration.
- `app/models` contains mocked business records.
- `app/structures` contains request/response schemas.

No database, Redis, RabbitMQ, tracing, middleware, config loader, or external
API integration is used.

## Run

```bash
go run .
```

Default port is `8088`. Override with `PORT=8090 go run .`.

## Selected APIs

| Method | Path | Why it is useful |
| --- | --- | --- |
| GET/POST | `/go/api/nsd/v1/fsa/GstSearch` | GST lookup with alternative search inputs, nested response, deprecated field, undocumented `source` param |
| GET/POST | `/go/api/nsd/v1/GstDetails/gstuserdetails` | GST registration/challan style schema with validation and nested business fields |
| GET | `/go/api/nsd/v1/userlist/companydetailnew` | Company profile, contacts, verification state, assignment, and paid service summary |
| GET | `/go/api/nsd/v1/NSDDashboard/MeetingScreenCheck` | Meeting eligibility workflow with nested counters and restriction state |
| GET | `/go/api/nsd/v1/tele/getVideoMeetData` | Video meeting status tied to GLID and meeting workflow |
| GET | `/go/api/nsd/v1/userlisting/mycompanylist` | Ranked company listing with filters, nested mobile list, priority/rank fields |

## Demo Drift Hooks

- Stale comments in controllers intentionally disagree with actual structs.
- `docs/stale_docs.md` intentionally documents outdated fields.
- `GSTSearchRequest.source` is accepted but absent from stale docs.
- `CompanyListItem.companyBucketName` and `GSTSearchData.legacyCompanyName`
  are deprecated response fields.
- Stale docs use `gst_company_name`, `meeting_allowed`, and `meetCode`, while
  actual responses use `gstCompanyName`, `isMeetingAllowed`, and
  `meeting_code`.

## Example

```bash
curl "http://localhost:8088/go/api/nsd/v1/fsa/GstSearch?empid=98123&AK=demo&gst=07AAECS1234F1Z5&lat=28.6&lng=77.2&source=mobile_app"
```
