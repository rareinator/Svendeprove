using DataAccessLibrary;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Configuration;
using System.Threading.Tasks;
using System.Net.Http;
using Blazored.Modal;
using System;
using Microsoft.AspNetCore.Components.WebAssembly.Authentication;
using Microsoft.AspNetCore.Components.Authorization;

namespace WebApp
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var builder = WebAssemblyHostBuilder.CreateDefault(args);
            builder.RootComponents.Add<App>("#app");

            builder.Services.AddScoped<CorsRequestAuthorizationMessageHandler>(); 
            builder.Services
                .AddHttpClient("BlazorClient.ServerApi", client => client.BaseAddress = new Uri(builder.Configuration.GetValue<string>("ServerApi:BaseAddress")))
                .AddHttpMessageHandler<CorsRequestAuthorizationMessageHandler>();

            builder.Services.AddScoped(sp => sp.GetRequiredService<IHttpClientFactory>().CreateClient("BlazorClient.ServerApi"));

            builder.Services.AddOidcAuthentication(options =>
            {
                options.ProviderOptions.Authority = builder.Configuration.GetValue<string>("Okta:Authority");
                options.ProviderOptions.ClientId = builder.Configuration.GetValue<string>("Okta:ClientId");

                options.ProviderOptions.ResponseType = "code";
                options.ProviderOptions.DefaultScopes.Add("profile");
                options.ProviderOptions.DefaultScopes.Add("address");
                options.ProviderOptions.DefaultScopes.Add("hospi");

                options.UserOptions.RoleClaim = "role";
            }).AddAccountClaimsPrincipalFactory<RolesClaimsPrincipalFactory>();

            builder.Services.AddApiAuthorization();


            builder.Services.AddScoped<IUserData, UserData>();
            builder.Services.AddScoped<IPatientData, PatientData>();
            builder.Services.AddScoped<IJournalData, JournalData>();
            builder.Services.AddScoped<IBookingData, BookingData>();

            builder.Services.AddBlazoredModal();
            await builder.Build().RunAsync();
        }
    }
}
