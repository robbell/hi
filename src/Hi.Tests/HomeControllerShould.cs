using System.Collections.Generic;
using FluentAssertions;
using Hi.Web;
using Hi.Web.Controllers;
using Microsoft.AspNetCore.Mvc;
using Xunit;

namespace Hi.Tests
{
    public class HomeControllerShould
    {
        [Fact]
        public void ListLatestPostsFromMedium()
        {
            var expectedPosts = new List<Post> { new Post { Title = "My Title", Body = "My Body" } };

            var controller = new HomeController();

            var result = (ViewResult)controller.Index();

            result.Model.Should().BeEquivalentTo(expectedPosts);
        }
    }
}
