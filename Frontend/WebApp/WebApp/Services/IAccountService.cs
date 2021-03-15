using DataAccessLibrary.Models;

namespace DataAccessLibrary
{
    public interface IAccountService
    {
        bool Login(UserModelOld user);
        bool Logout();
    }
}