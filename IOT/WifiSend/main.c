#define F_CPU 16000000UL			/* Define CPU Frequency - 16MHz */
#include <avr/io.h>
#include <util/delay.h>
#include <stdbool.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <avr/interrupt.h>
#include <avr/wdt.h>
#include "USART/USART_RS232_H_file.h"
#include "DHT/DHT.h"
#include "ESP8266/ESP8266.h"

uint8_t sensorData[5];
uint8_t sensorReadSuccess = 0;

ISR (USART1_RX_vect)
{
	uint8_t oldsrg = SREG;
	cli();
	RESPONSE_BUFFER[Counter] = UDR1;
	Counter++;
	if(Counter == DEFAULT_BUFFER_SIZE){
		Counter = 0; pointer = 0;
	}
	SREG = oldsrg;
}

ISR(TIMER3_COMPA_vect)
{
	sensorReadSuccess = DHTSensorRead(sensorData);
	wdt_reset();	// Reset watchdog timer on every sensor read
}



int main(void)
{
	USART_Init(115200);						/* Initiate USART with 115200 baud rate */
	DHTInit();								/* Setup timer used for DHT */
	ESP8266_Init();							/* Setup module and connect to Wifi */

	while(1)
	{
		// Switch between temperature and humidity sensor data
		for (uint8_t i = 0; i < 2; i++)
		{
			_delay_ms(2000);
			char sensor[10];
			uint8_t sensorRead;
			if (i == 0)
			{
				sensorRead = 0;
				sprintf(sensor, "dht-hum");
			}
			else
			{
				sensorRead = 2;
				sprintf(sensor, "dht-temp");
			}
			if (sensorReadSuccess) // Only send data to the API if the sensor has been read.
			SendApiData(sensor, sensorRead, sensorData);
		}
	}
}
