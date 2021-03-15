﻿using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IBookingData
    {
        Task<List<BookingModel>> GetBookingsByPatient(string patient);
        void DeleteBooking(int bookingId);
        Task<BookingModel> InsertBooking(BookingModel booking);
        Task<List<HospitalModel>> GetHospitals();
        Task<List<UserModel>> GetAvailableEmployees(int hospitalId);
        Task<List<DateTime>> GetAvailableTimes(AvailableEmpDto data);
    }
}