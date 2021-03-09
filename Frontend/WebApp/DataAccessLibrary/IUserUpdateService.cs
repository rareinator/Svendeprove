using DataAccessLibrary.Models;
using System;

namespace DataAccessLibrary
{
    public interface IUserUpdateService
    {   
        event Action<UserModel> OnUserUpdate;
        void UpdateUser(UserModel user);
    }
}