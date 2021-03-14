using DataAccessLibrary.Models;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IUserData
    {
        Task<TokenModel> Login(UserModel user, string scope);
    }
}