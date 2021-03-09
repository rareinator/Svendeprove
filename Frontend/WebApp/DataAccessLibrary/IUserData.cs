using DataAccessLibrary.Models;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IUserData
    {
        Task<string> Login(UserModel user);
    }
}