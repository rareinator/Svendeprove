using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text;
using System.Threading.Tasks;
using System.Text.Json;

namespace DataAccessLibrary
{

    public class TokenClass
    {
        public string Token {get; set;}
    }


    public class UserData : IUserData
    {
        private readonly HTTPService _client;

        public UserData(HTTPService client)
        {
            _client = client;
        }

        public async Task<string> Login(UserModel user)
        {
            var token = await _client.PostData<UserModel,TokenClass>("/authentication/patient/login", user);

            return token.Token;
        }
    }
}
