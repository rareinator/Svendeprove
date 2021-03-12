using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Json;
using System.Net.Http.Headers;
using System.Text;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public class HTTPService
    {
        private readonly HttpClient _client;

        public HTTPService(Uri baseAddress)
        {
          _client = new HttpClient {BaseAddress = baseAddress};
        }

        public async Task<U> PostData<T,U>(string uri, T data)
        {
          if (_client.DefaultRequestHeaders.Contains("Authorization"))
          {
            var response = await _client.PostAsJsonAsync<T>(uri, data);

            if (response.IsSuccessStatusCode)
            {
              return await response.Content.ReadFromJsonAsync<U>();
            }
          }
          return default(U);
        }

        public async Task<T> GetData<T>(string uri)
        {
          if (_client.DefaultRequestHeaders.Contains("Authorization"))
          {
            return await _client.GetFromJsonAsync<T>(uri);
          }
          return default(T);
        }

        public async Task SetAuthHeader(AuthenticationHeaderValue header)
        {
          if (_client.DefaultRequestHeaders.Contains("Authorization"))
          {
              _client.DefaultRequestHeaders.Remove("Authorization");
          }
          _client.DefaultRequestHeaders.Authorization = header;
        }
    }
}
