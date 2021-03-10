using DataAccessLibrary;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Blazored.Modal;

namespace WebApp
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var builder = WebAssemblyHostBuilder.CreateDefault(args);
            builder.RootComponents.Add<App>("app");

            builder.Services.AddScoped(sp => new HttpClient { BaseAddress = new Uri("http://cloud.m9ssen.me:56060") });
            builder.Services.AddScoped<IUserData, UserData>();
            builder.Services.AddSingleton<IUserUpdateService, UserUpdateService>();
            builder.Services.AddScoped<IPatientData, PatientData>();
            builder.Services.AddScoped<IJournalData, JournalData>();

            builder.Services.AddBlazoredModal();

            await builder.Build().RunAsync();
        }
    }
}
