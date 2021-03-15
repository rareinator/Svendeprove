using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public class BookingData : IBookingData
    {
        private readonly HttpClient _client;

        public BookingData(HttpClient client)
        {
            _client = client;
        }


        public async Task<List<BookingModel>> GetBookingsByPatient(string patient)
        {
            return await _client.GetFromJsonAsync<List<BookingModel>>($"/booking/byPatient/{patient}");
        }

        public async void DeleteBooking(int bookingId)
        {
            await _client.DeleteAsync($"/booking/{bookingId}");
        }

        public async Task<BookingModel> InsertJournal(BookingModel booking)
        {
            var response = await _client.PostAsJsonAsync($"/booking", booking);
            string responseMessage = await response.Content.ReadAsStringAsync();

            BookingModel responseJournal = JsonSerializer.Deserialize<BookingModel>(responseMessage);

            return responseJournal;
        }
    }
}
