const byte UP_INPUT = 7;
const byte RIGHT_INPUT = 6;
const byte DOWN_INPUT = 5;
const byte LEFT_INPUT = 4;

void pciSetup(byte pin)
{
    *digitalPinToPCMSK(pin) |= bit (digitalPinToPCMSKbit(pin));  // enable pin
    PCIFR  |= bit (digitalPinToPCICRbit(pin)); // clear any outstanding interrupt
    PCICR  |= bit (digitalPinToPCICRbit(pin)); // enable interrupt for the group
}
void setup(){
  Serial.begin(115200);
  pciSetup(UP_INPUT);
  pciSetup(RIGHT_INPUT);
  pciSetup(DOWN_INPUT);
  pciSetup(LEFT_INPUT);
}

ISR (PCINT2_vect) // handle pin change interrupt for D0 to D7 here
{
  if (digitalRead(UP_INPUT) == HIGH)
    Serial.write(UP_INPUT);
  else if (digitalRead(RIGHT_INPUT) == HIGH)
    Serial.write(RIGHT_INPUT);
  else if (digitalRead(DOWN_INPUT) == HIGH)
    Serial.write(DOWN_INPUT);
  else if (digitalRead(LEFT_INPUT) == HIGH)
    Serial.write(LEFT_INPUT);
}

void loop() {
  // Nothing needed
}
