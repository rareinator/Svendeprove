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
        private readonly HttpClient _client;

        public UserData(HttpClient client)
        {
            _client = client;
        }

        public async Task<string> Login(UserModel user)
        {
            var response = await _client.PostAsJsonAsync<UserModel>("/authentication/patient/login", user);
            string responseMessage = await response.Content.ReadAsStringAsync();

            var token = JsonSerializer.Deserialize<TokenClass>(responseMessage);



            return token.Token;
        }
    }
}
