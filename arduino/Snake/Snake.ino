void setup(){
  Serial.begin(115200);
}

void loop(){
  byte b = 65;
  Serial.write(b);
  delay(1000);
}
