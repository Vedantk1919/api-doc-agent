# Intentionally Stale API Notes

This file is deliberately wrong in a few places so API DocAgent can demonstrate
drift detection between old documentation and live Go structs/routes.

## GET /go/api/nsd/v1/fsa/GstSearch

Old request docs:

| Field | Required | Notes |
| --- | --- | --- |
| empid | yes | Employee id |
| AK | yes | Auth token |
| gst | yes | GSTIN |

Known drift:

- The real API also accepts `email`, `mobile`, `lat`, `lng`, `MODID`, and the undocumented `source` parameter.
- Old docs say the response contains `data.gst_company_name`.
- Actual response contains `data.gstCompany.gstCompanyName`.
- Actual response includes deprecated `data.legacyCompanyName`.

## GET /go/api/nsd/v1/GstDetails/gstuserdetails

Old request docs omit `duplicateFlag`, `check_master_table`, `fcprealtime`,
`platform`, and `os`.

Old response example:

```json
{
  "status": 200,
  "message": "Data Found",
  "data": {
    "gst_status": "ACTIVE"
  }
}
```

Actual response nests GST registration data under `data.result.gstin_status`.

## GET /go/api/nsd/v1/userlist/companydetailnew

Old docs say `companyAddress` is top-level. Actual response nests location as
`data.company.city`, `data.company.state`, and `data.company.pincode`.

## GET /go/api/nsd/v1/NSDDashboard/MeetingScreenCheck

Old docs say the response field is `meeting_allowed`. Actual response uses
`data.isMeetingAllowed`.

The `newflow` request flag is intentionally missing from this old document.

## GET /go/api/nsd/v1/tele/getVideoMeetData

Old docs call the response field `meetCode`. Actual response uses
`data.meeting_code`.

## GET /go/api/nsd/v1/userlisting/mycompanylist

Old request docs omit `risk_bucket`.

Old docs say company records contain `bucketName`. Actual response keeps a
deprecated compatibility field called `companyBucketName` and uses `rank` plus
`priorityKeyStr` for current ranking.
