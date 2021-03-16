
using System.Linq;
using System.Security.Claims;
using System.Text.Json;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components.WebAssembly.Authentication;
using Microsoft.AspNetCore.Components.WebAssembly.Authentication.Internal;

namespace WebApp
{
  public class RolesClaimsPrincipalFactory : AccountClaimsPrincipalFactory<RemoteUserAccount>
  {
      public RolesClaimsPrincipalFactory(IAccessTokenProviderAccessor accessor) : base(accessor)
      {
      }

      public override async ValueTask<ClaimsPrincipal> CreateUserAsync(
          RemoteUserAccount account, RemoteAuthenticationUserOptions options)
      {
          var user = await base.CreateUserAsync(account, options);
          if (!user.Identity.IsAuthenticated)
          {
              return user;
          }

          var identity = (ClaimsIdentity) user.Identity;
          var roleClaims = identity.FindAll(claim => claim.Type == "groups").ToList();
          if (roleClaims == null || !roleClaims.Any())
          {
              return user;
          }

          foreach (var existingClaim in roleClaims)
          {
              identity.RemoveClaim(existingClaim);
          }

          var rolesElem = account.AdditionalProperties["groups"];
          if (!(rolesElem is JsonElement roles))
          {
              return user;
          }

          if (roles.ValueKind == JsonValueKind.Array)
          {
              foreach (var role in roles.EnumerateArray())
              {
                  identity.AddClaim(new Claim(options.RoleClaim, role.GetString()));
              }
          }
          else
          {
              identity.AddClaim(new Claim(options.RoleClaim, roles.GetString()));
          }

          return user;
      }
  }

}
