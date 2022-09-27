#include <M5Stack.h>
#include <WiFi.h>
#include <WiFiClient.h>

#include "secrets.h"

void connectWiFi() {
  WiFi.begin(SSID, PASSPHRASE);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
}



void setup() {
  M5.begin();
  M5.Power.begin();
  M5.Lcd.setTextSize(3);
  M5.Lcd.print("starting...\n");
  M5.Lcd.print("connecting to wifi...\n");
  connectWiFi();
}

void loop() {
  M5.update();
  if (WiFi.status() == WL_CONNECTED) {
    M5.Lcd.printf("\nConnected!\n");
    M5.Lcd.printf("IP address: %s\n", WiFi.localIP().toString().c_str());
  }
  if (M5.BtnA.pressedFor(1000, 200)) { // If A is pressed for 1 second power off
    M5.Power.powerOFF();
  }
}
