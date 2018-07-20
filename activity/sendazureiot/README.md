# Send to azure IOTHub
This activity provides your Flogo app the ability to send message to the Azure Iot Hub from a device

## Installation

```bash
flogo install github.com/pradyuz3rocool/flogo-workspace/activity/sendazureiot
```
Link for flogo web:
```
https://github.com/pradyuz3rocool/flogo-workspace/activity/sendazureiot
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
      "name": "connectionString",
      "type": "string",
      "required": true
    },
    {
      "name": "Device ID",
      "type": "string"
    },
    {
      "name": "Action",
      "type": "string",
      "required": true,
      "allowed": ["Send"]
    },
    {
      "name": "message",
      "type": "string"
    }
  ]
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| Connection String              | Your Azure IOT Connection String.            |
| Device ID                      | Name of the Device  |
| Action                         | Send                |
| Mesage                         | Message to be sent  |

## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| result       | A response from the request, It would be something like `'{"deviceId":"testing","generationId":"636595531817773533","etag":"NzkxOTM4Njcx","connectionState":"Disconnected","status":"enabled","statusReason":null,"connectionStateUpdatedTime":"0001-01-01T00:00:00","statusUpdatedTime":"0001-01-01T00:00:00","lastActivityTime":"0001-01-01T00:00:00","cloudToDeviceMessageCount":0,"authentication":{"symmetricKey":{"primaryKey":"**********","secondaryKey":"**************"},"x509Thumbprint":{"primaryThumbprint":null,"secondaryThumbprint":null}}}'` |
| status       | The status of the request made                            |