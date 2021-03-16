using System;
using System.Collections.Generic;
using System.Text;

namespace DataAccessLibrary.Models
{
    public class UserModelOld
    {
        public string Username { get; set; }
        public string Password { get; set; }
        public string Token { get; set; }
        public string Role { get; set; }

        public string Scope { get; set; }
        public int UserID { get; set; }
        public string Name { get; set; }
    }
}
