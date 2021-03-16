using DataAccessLibrary.Converters;
using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Linq;
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

        public async Task<List<BookingModel>> GetBookingsByEmployee(string employee)
        {
            return await _client.GetFromJsonAsync<List<BookingModel>>($"/booking/byEmployee/{employee}");
        }

        public async void DeleteBooking(int bookingId)
        {
            await _client.DeleteAsync($"/booking/{bookingId}");
        }

        public async Task<BookingModel> InsertBooking(BookingModel booking)
        {
            var response = await _client.PostAsJsonAsync($"/booking", booking);
            string responseMessage = await response.Content.ReadAsStringAsync();

            BookingModel responseJournal = JsonSerializer.Deserialize<BookingModel>(responseMessage);

            return responseJournal;
        }

        public async Task<BookingModel> UpdateBooking(BookingModel booking)
        {
            var response = await _client.PostAsJsonAsync($"/booking/{booking.BookingId}", booking);
            string responseMessage = await response.Content.ReadAsStringAsync();

            BookingModel responseJournal = JsonSerializer.Deserialize<BookingModel>(responseMessage);

            return responseJournal;
        }

        public async Task<List<HospitalModel>> GetHospitals()
        {
            return await _client.GetFromJsonAsync<List<HospitalModel>>($"/admin/hospitals");
        }

        public async Task<List<UserModel>> GetAvailableEmployees(int hospitalId)
        {
            return await _client.GetFromJsonAsync<List<UserModel>>($"/admin/doctors/inHospital/{hospitalId}");
        }

        public async Task<List<DateTime>> GetAvailableTimes(AvailableEmpDto data)
        {

            var response = await _client.PostAsJsonAsync($"/booking/availableTimesForDoctor", data);
            string responseMessage = await response.Content.ReadAsStringAsync();

            var options = new JsonSerializerOptions
            {
                Converters =
                {
                    new DateTimeConverter()
                }
            };
            List<DateTime> responseTimes = JsonSerializer.Deserialize<List<DateTime>>(responseMessage, options);

            return responseTimes;
        }

        public async Task<List<DepartmentModel>> GetDepartments()
        {
            return await _client.GetFromJsonAsync<List<DepartmentModel>>($"/admin/departments");
        }

        public async Task<List<DepartmentModel>> GetAvailableDepartments(int hospitalId)
        {
            List<DepartmentModel> departments = await GetDepartments();

            return departments.Where(x => x.HospitalId == hospitalId).ToList();
        }
        
        public async Task<List<BedModel>> GetBeds()
        {
            return await _client.GetFromJsonAsync<List<BedModel>>($"/admin/beds");
        }

        public async Task<List<BedModel>> GetAvailableBeds(int departmentId)
        {
            List<BedModel> beds = await GetBeds();

            return beds.Where(x => x.DepartmentId == departmentId && x.IsAvailable == true).ToList();
        }
    }
}
