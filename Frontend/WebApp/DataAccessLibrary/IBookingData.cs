using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IBookingData
    {
        Task<List<BookingModel>> GetBookingsByPatient(string patient);
        Task<List<BookingModel>> GetBookingsByEmployee(string employee);
        void DeleteBooking(int bookingId);
        Task<BookingModel> InsertBooking(BookingModel booking);
        Task<BookingModel> UpdateBooking(BookingModel booking);
        Task<List<HospitalModel>> GetHospitals();
        Task<List<UserModel>> GetAvailableEmployees(int hospitalId);
        Task<List<DateTime>> GetAvailableTimes(AvailableEmpDto data);
        Task<List<DepartmentModel>> GetDepartments();
        Task<List<DepartmentModel>> GetAvailableDepartments(int hospitalId);
        Task<List<BedModel>> GetBeds();
        Task<List<BedModel>> GetAvailableBeds(int departmentId);
    }
}