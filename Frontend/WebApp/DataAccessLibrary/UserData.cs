using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text;
using System.Threading.Tasks;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace DataAccessLibrary
{

    public class CodeClass
    {
        [JsonPropertyName("code")]
        public string Code {get; set;}
    }

    public class TokenModel
    {

        [JsonPropertyName("expires_in")]
        public int ExpiresIn { get; set; }
        
        [JsonPropertyName("access_token")]
        public string AccessToken { get; set; }

        [JsonPropertyName("scope")]
        public string Scope { get; set; }

        [JsonPropertyName("token_type")]
        public string TokenType { get; set; }

    }


    public class UserData : IUserData
    {
        private readonly HttpClient _client;

        public UserData(HttpClient client)
        {
            _client = client;
        }

        public async Task<TokenModel> Login(UserModel user, string scope)
        {
            var codeResponse = await _client.PostAsJsonAsync<UserModel>(string.Format("/auth?response_type=code&client_id=000000&scope={0}",scope), user);
            string codeResponseMessage = await codeResponse.Content.ReadAsStringAsync();

            var code = JsonSerializer.Deserialize<CodeClass>(codeResponseMessage);
            
            var tokenResponse = await _client.PostAsJsonAsync<CodeClass>("/token",code);
            string tokenResponseMessage = await tokenResponse.Content.ReadAsStringAsync();

            var token = JsonSerializer.Deserialize<TokenModel>(tokenResponseMessage);

            return token;
        }
    }
}
