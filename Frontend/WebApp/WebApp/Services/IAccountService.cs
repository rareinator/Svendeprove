using DataAccessLibrary.Models;

namespace DataAccessLibrary
{
    public interface IAccountService
    {
        bool Login(UserModel user);
        bool Logout();
    }
}