#include <M5Stack.h>
#include <WiFi.h>
#include <WiFiClient.h>
#include <WiFiUdp.h>

#include "secrets.h"

void connectWiFi() {
  WiFi.begin(SSID, PASSPHRASE);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
}

WiFiUDP udp;

void sendUDP(uint8_t *buffer, size_t size) {
  IPAddress broadcastIp(192,168,0,255);
  udp.beginPacket(broadcastIp, 8899); 
  udp.write(buffer, size);
  udp.endPacket();
}

void readUDP() {
  // udp.begin(8899);
  String req;
  if (udp.parsePacket() > 0) {
    req = "";
    while (udp.available() > 0) {
      char z = udp.read();
      req += z;
    }
    M5.Lcd.print(req);
  }

  // udp.stop();
}

void setup() {
  M5.begin();
  M5.Power.begin();

  M5.Lcd.setTextColor(ORANGE, BLACK);
  M5.Lcd.setTextSize(2.5);

  M5.Lcd.print("starting...\n");
  M5.Lcd.print("connecting to wifi...\n");
  connectWiFi();
  udp.begin(8899);
}

void loop() {
  static unsigned long previousMillis = 0;
  unsigned long currentMillis = millis();
  unsigned long delta = currentMillis - previousMillis;
  M5.update();

  M5.Lcd.setCursor(0, 0);

  if (WiFi.status() == WL_CONNECTED) {
    M5.Lcd.printf("Connected! \n");
    M5.Lcd.printf("IP address: %s\n", WiFi.localIP().toString().c_str());
  }

  readUDP();

  if (M5.BtnA.wasReleased()) {
    static bool isOn = false;
    M5.Lcd.invertDisplay(isOn);
    isOn = !isOn;
  }
  if (M5.BtnA.pressedFor(1000, 200)) { // If A is pressed for 1 second power off
    M5.Power.powerOFF();
  }
  if (M5.BtnB.wasReleased()) {
    sendUDP((uint8_t *)"B", 1);  
  }
  if (M5.BtnC.wasReleased()) {
    sendUDP((uint8_t *)"C", 1);  
  }


}
