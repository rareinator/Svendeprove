using DataAccessLibrary.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IBookingData
    {
        Task<List<BookingModel>> GetBookingsByPatient(string patient);
        void DeleteBooking(int bookingId);
        Task<BookingModel> InsertJournal(BookingModel booking);
    }
}