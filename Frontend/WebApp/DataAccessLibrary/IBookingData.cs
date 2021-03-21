using DataAccessLibrary.Models;
using DataAccessLibrary.TransferObjects;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IBookingData
    {
        /// <summary>
        /// Sends a GET request to enpoint "/booking/byPatient/{param} to get bookings by the specified patient 
        /// </summary>
        /// <param name="patient"></param>
        /// <returns>List of patient booking objects</returns>
        Task<List<BookingModel>> GetBookingsByPatient(string patient);

        /// <summary>
        /// Sends a GET request to enpoint "/booking/byEmployee/{param} to get bookings by the specified employee attending in the booking 
        /// </summary>
        /// <param name="employee"></param>
        /// <returns>List of patient booking objects</returns>
        Task<List<BookingModel>> GetBookingsByEmployee(string employee);

        /// <summary>
        /// Sends a DELETE request to endpoint "/booking/{param}" to delete the booking with the specified documentId
        /// </summary>
        /// <param name="bookingId"></param>
        Task DeleteBooking(int bookingId);

        /// <summary>
        /// Sends a POST request to endpoint "/booking with {param} in the request body to add a new patient booking
        /// </summary>
        /// <param name="booking"></param>
        /// <returns>New added booking object</returns>
        Task<BookingModel> InsertBooking(BookingModel booking);

        /// <summary>
        /// Sends a POST request to endpoint "/booking/{param.BookingId} with {param} in the request body to edit a existing booking
        /// </summary>
        /// <param name="booking"></param>
        /// <returns>Edited booking object</returns>
        Task<BookingModel> UpdateBooking(BookingModel booking);

        /// <summary>
        /// Sends a GET request to enpoint "/admin/hospitals to get all hospitals
        /// </summary>
        /// <returns>List of hospital objects</returns>
        Task<List<HospitalModel>> GetHospitals();

        /// <summary>
        /// Sends a GET request to enpoint "/admin/doctors/inHospital/{param} to get available employees in specified hospital by id
        /// </summary>
        /// <param name="hospitalId"></param>
        /// <returns>List of user objects</returns>
        Task<List<UserModel>> GetAvailableEmployees(int hospitalId);

        /// <summary>
        /// Sends a POST request to enpoint "/booking/availableTimesForDoctor with {param} in the request body to get all available times for specified employee and day
        /// </summary>
        /// <param name="data"></param>
        /// <returns>List of datetimes</returns>
        Task<List<DateTime>> GetAvailableTimes(AvailableEmpDto data);

        /// <summary>
        /// Sends a POST request to enpoint "/admin/availableBeds with {param} in the request body to get all available beds for specified hospitalId and timeframe
        /// </summary>
        /// <param name="data"></param>
        /// <returns>List of bed objects</returns>
        Task<List<BedModel>> GetAvailableBeds(AvailableBedDto data);
    }
}