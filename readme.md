# OpenPrio Proxy

OpenPrio Proxy is a tool that can help transport operators to start share data with the OpenPrio message format. Sometimes it's difficult to add modern techniques like protocol buffers and mqtt clients into existing software, OpenPrio Proxy is created to support those systems. It should be run in an internal network and creates a bridge between the internal network and a public MQTT-broker. Data can be provided over a TCP-socket in CSV or protocol buffer (if your software support protocol buffers this is the preferred method) format.

## Setup
OpenPrio proxy is setup by environment variablefs, in this repository an example (.env_example) provided. All MQTT variables are required, the other variables are optional and have default values. 

| Variable               | Description                                                                                                    | Default value     |
| -----------------------| -------------------------------------------------------------------------------------------------------------- | -----------------:|
| DATA_FORMAT_OF_SOURCE  | Describes in what format the data is provided to the proxy, possible values are 'csv' and 'proto'              | 'csv'             |
| TCP_LISTEN_ADDRESS     | The TCP address the proxy listens, this should be a address within an internal network.                        | '127.0.0.1:9001'  |
| MQTT_HOST              | The host of the MQTT broker the data should be send to. (be aware only tls/ssl connections are supported)      |  None             |
| MQTT_USERNAME          | The username for MQTT.                                                                                         |  None             |
| MQTT_PASSWORD          | The MQTT password.                                                                                             |  None             |

## Format

### CSV

Expected order of fields:
data_ownercode, block_code, vehicle_number, latitude (wsg84), longitude (wsg84), accuracy (m), speed (m/s), bearing (degrees), odometer (m), hdop, timestamp ( milliseconds since epoch), dooropeningstatus (see .proto), stopbuttonstatus (see .proto)\n
Example:
QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,1613652529565,0,1\n


