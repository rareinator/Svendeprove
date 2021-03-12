using DataAccessLibrary.Models;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IUserData
    {
        Task<string> LoginPatient(UserModel user);
        Task<string> LoginEmployee(UserModel user);
    }
}