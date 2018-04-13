using System.Collections.Generic;
using System.Diagnostics;
using Microsoft.AspNetCore.Mvc;
using Hi.Web.Models;

namespace Hi.Web.Controllers
{
    public class HomeController : Controller
    {
        public IActionResult Index()
        {
            return View(new List<Post>{new Post { Title = "My Title", Body = "My Body" }});
        }

        public IActionResult About()
        {
            ViewData["Message"] = "Your application description page.";

            return View();
        }

        public IActionResult Contact()
        {
            ViewData["Message"] = "Your contact page.";

            return View();
        }

        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }
    }
}
