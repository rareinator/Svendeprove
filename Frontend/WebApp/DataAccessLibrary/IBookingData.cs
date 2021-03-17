using DataAccessLibrary.Models;
using DataAccessLibrary.TransferObjects;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IBookingData
    {
        Task<List<BookingModel>> GetBookingsByPatient(string patient);
        Task<List<BookingModel>> GetBookingsByEmployee(string employee);
        Task DeleteBooking(int bookingId);
        Task<BookingModel> InsertBooking(BookingModel booking);
        Task<BookingModel> UpdateBooking(BookingModel booking);
        Task<List<HospitalModel>> GetHospitals();
        Task<List<UserModel>> GetAvailableEmployees(int hospitalId);
        Task<List<DateTime>> GetAvailableTimes(AvailableEmpDto data);
        Task<List<BedModel>> GetAvailableBeds(AvailableBedDto data);
    }
}