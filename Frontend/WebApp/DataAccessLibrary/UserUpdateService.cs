using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Text;

namespace DataAccessLibrary
{
    public class UserUpdateService : IUserUpdateService
    {
        public event Action<UserModel> OnUserUpdate;

        public void UpdateUser(UserModel user)
        {
            //OnUserUpdate?.Invoke(user);
        }
    }
}
