#include "measurements.pb.h"
#include <WiFiManager.h> 
#include <PubSubClient.h>
#include <WiFiClientSecure.h>
#include <BME280I2C.h>
#include "pb_common.h"
#include "pb_encode.h"
#include "pb_decode.h"
#include <Wire.h>
#include "pb.h"
#include "conf.h"

#define SERIAL_BAUD 115200
#define SAMPLING_SIZE 100
#define SLEEP_TIME 3.6e9 // 1h //3e8; //5 min

int chipID = ESP.getChipId();
String pubTopic = String(String("terrasense/")+chipID+String("/measurements"));
float soilMoistureValue(NAN);
float temp(NAN), hum(NAN), pres(NAN);
BME280::TempUnit tempUnit(BME280::TempUnit_Celsius);
BME280::PresUnit presUnit(BME280::PresUnit_Pa);

/**** Secure WiFi Connectivity Initialisation *****/
WiFiClientSecure espClient;

/**** MQTT Client Initialisation Using WiFi Connection *****/
PubSubClient client(espClient);

WiFiManager wifiManager;
BME280I2C bme;  // Default : forced mode, standby time = 1000 ms
                // Oversampling = pressure ×1, temperature ×1, humidity ×1, filter off,

/************* Connect to MQTT Broker ***********/
void reconnect() {
  // Loop until we're reconnected
  while (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    // Attempt to connect
    if (client.connect(String(chipID).c_str(), MQTT_USERNAME, MQTT_PASSWORD)) {
      Serial.println("connected");

      client.subscribe("led_state");   // subscribe the topics here

    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");   // Wait 5 seconds before retrying
      delay(5000);
    }
  }
}

pb_Measurements createMeasurement(float soil, float temp, float pres, float hum){
  pb_Measurements m = pb_Measurements_init_zero;
  m.chipID = chipID;
  m.soilMoisture = soil;
  m.temperature = temp;
  m.pressure = pres;
  m.humidity = hum;
  return m;
}

void sendMeasurement(const char* topic, pb_Measurements m, boolean retained){
  uint8_t buffer[100];
  auto inputLength = sizeof(buffer);
  pb_ostream_t stream = pb_ostream_from_buffer(buffer, inputLength);
  
  if (!pb_encode(&stream, pb_Measurements_fields, &m)){
    Serial.println("failed to encode proto");
    return;
  }
  Serial.print("This is the number of bytes written: ");
  Serial.println(stream.bytes_written);

  if(client.publish(topic, buffer, stream.bytes_written, retained)){
    Serial.println("Message publised ["+String(topic)+"]");
  }
}

//////////////////////////////////////////////////////////////////
void setup() {
  Serial.begin(SERIAL_BAUD);

  while (!Serial) {}  // Wait

  Wire.begin();

  while (!bme.begin()) {
    Serial.println("Could not find BME280 sensor!");
  }
  wifiManager.autoConnect("TerraSense_AP");

  espClient.setInsecure();
  client.setServer(MQTT_SERVER, MQTT_PORT);
}

//////////////////////////////////////////////////////////////////
void loop() {
  delay(100);
  Serial.println("Entering the loop");
  if (!client.connected()) reconnect();
  int sum = 0;
  for(int i=0; i<SAMPLING_SIZE; i++){
    sum += analogRead(A0);
    delay(10);
  }
  soilMoistureValue = sum/SAMPLING_SIZE;
  bme.read(pres, temp, hum, tempUnit, presUnit);
  Serial.println("create measurement");
  pb_Measurements measurements = createMeasurement(soilMoistureValue, temp, pres, hum);
  sendMeasurement(pubTopic.c_str(), measurements, false);
  Serial.println("now let's take a nap...");
  // ESP.deepSleep(SLEEP_TIME); //deep sleep
  delay(4e6); //delay - in case of powerbank connection - 2 hours
}