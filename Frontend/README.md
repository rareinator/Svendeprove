# Architechture

### Projects

* WebApp (Blazor Web Assembly Web App - ASP.NET Core 5.0)
* DataAccessLibrary (.NET Standar 2.0)

### Project References
* WebApp -> DataAccesLibrary

### Packages

```powershell
Install-Package Blazored.Typeahead -ProjectName WebApp
Install-Package Blazored.Modal -ProjectName WebApp
Install-Package Blazored.TextEditor -ProjectName WebApp
Install-Package Microsoft.Extensions.Http -ProjectName WebApp
Install-Package System.Net.Http.Json -ProjectName WebApp
Install-Package System.Net.Http.Json -ProjectName DataAccesLibrary
```

[Blazored Typeahead](https://github.com/Blazored/Typeahead)

[Blazored Modal](https://github.com/Blazored/Modal)

[Blazored TextEditor](https://github.com/Blazored/TextEditor)

[Http Extensions](https://www.nuget.org/packages/Microsoft.Extensions.Http/5.0.0)

[Http JSON](https://www.nuget.org/packages/System.Net.Http.Json/5.0.0)

# Dependency Injection
All dependency injections  should be kept in the Program.cs file and is done as follows:

```C#
builder.Services.AddScoped<IPatientData, PatientData>();
builder.Services.AddScoped<IJournalData, JournalData>();
builder.Services.AddScoped<IBookingData, BookingData>();
```

# OAuth & OIDC Setup
Program.cs
```C#
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
```

There is 5 different role types from the Okta auhtentication/Authorization service:
* Office
* Doctor
* Nurse
* Patient

#### Usage
```xml
<AuthorizeView Roles="Doctor, Nurse">
    <Authorized>
        <p>Only doctors and nurses are authorized to see this</p>
    </Authorized
    <NotAuthorized>
        <p>If not a doctor or nurse then see this</p>
    <NotAuthorized>
</AuthorizeView>
```


