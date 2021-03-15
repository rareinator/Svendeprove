using DataAccessLibrary.Models;
using Microsoft.AspNetCore.Components.Authorization;
using System;
using System.Collections.Generic;
using System.Text;
using WebApp.Auth;

namespace DataAccessLibrary
{
    public class AccountService : IAccountService
    {
        private readonly AuthenticationStateProvider _authenticationStateProvider;
        public AccountService(AuthenticationStateProvider authenticationStateProvider)
        {
            _authenticationStateProvider = authenticationStateProvider;
        }
        public bool Login(UserModelOld user)
        {
            (_authenticationStateProvider as CustomAuthenticationProvider).LoginNotify(user);
            return true;
        }

        public bool Logout()
        {
            (_authenticationStateProvider as CustomAuthenticationProvider).LogoutNotify();
            return true;
        }
    }
}
