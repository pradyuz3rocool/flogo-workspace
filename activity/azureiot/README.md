# Azure IOT Hub Connector
This activity provides your Flogo app the ability to perform operations on the Azure Iot Hub

## Installation

```bash
flogo install github.com/pradyuz3rocool/flogo-workspace/activity/azureiot
```
Link for flogo web:
```
https://github.com/pradyuz3rocool/flogo-workspace/activity/azureiot
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
      "name": "connectionString",
      "type": "string",
      "required": true
    }
  ],
    "outputs": [ 
    {
      "name": "result",
      "type":"any"
    },
    {
      "name": "status",
      "type":"any"
    }
  ]
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| Connection String               | Your Azure IOT ConectionString.            |

## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| result       | A response from the request, It would be something like `'{"deviceId":"testing","generationId":"636595531817773533","etag":"NzkxOTM4Njcx","connectionState":"Disconnected","status":"enabled","statusReason":null,"connectionStateUpdatedTime":"0001-01-01T00:00:00","statusUpdatedTime":"0001-01-01T00:00:00","lastActivityTime":"0001-01-01T00:00:00","cloudToDeviceMessageCount":0,"authentication":{"symmetricKey":{"primaryKey":"**********","secondaryKey":"**************"},"x509Thumbprint":{"primaryThumbprint":null,"secondaryThumbprint":null}}}'` |
| status       | The status of the request made                            |