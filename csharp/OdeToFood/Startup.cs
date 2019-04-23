using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Routing;
using Microsoft.AspNetCore.Http;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using OdeToFood.Data;
using OdeToFood.Services;

namespace OdeToFood
{
    public class Startup
    {
        // Do this to get access to the configuration in ConfigureServices,
        // since that isn't an injectable method like Configure
        public Startup(IConfiguration configuration)
        {
            _configuration = configuration;
        }
        // This method gets called by the runtime. Use this method to add services to the container.
        // For more information on how to configure your application, visit https://go.microsoft.com/fwlink/?LinkID=398940
        public void ConfigureServices(IServiceCollection services)
        {
            /*
                services.
                    AddSingleton<>() - a single instance of the service
                    AddTransient<>() - a new instance each *use*
                    AddScoped<>() - a new instance each *request*
             */
            services.AddDbContext<OdeToFoodDBContext>(options =>
                options.UseSqlServer(_configuration.GetConnectionString("OdeToFood")));
            services.AddSingleton<IGreeter, Greeter>();
            // services.AddSingleton<IRestaurantData, InMemoryRestaurantData>();
            services.AddScoped<IRestaurantData, SqlRestaurantData>();
            services.AddMvc();
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app,
                              IHostingEnvironment env,
                              IGreeter greeter,
                              ILogger<Startup> logger)
        {
            // Default env var is ASPNETCORE_ENVIRONMENT
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }
            else
            {
                app.UseExceptionHandler();
            }

            // Serve static files, and use defaults like index.html
            // app.UseDefaultFiles();
            app.UseStaticFiles();
            // or app.UseFileServer() does both

            // Use MVC
            //app.UseMvcWithDefaultRoute();
            app.UseMvc(ConfigureRoutes);

            /*
            // Custom middleware
            app.Use(next => {
                return async context =>
                {
                    logger.LogInformation("Request incoming");
                    if(context.Request.Path.StartsWithSegments("/mym")) {
                        await context.Response.WriteAsync("Hit");
                        logger.LogInformation("Request handled");
                    } else {
                        await next(context);
                        // This is control flow going backward - after the rest of the middleware have executed
                        // and come back
                        logger.LogInformation("Request outgoing");
                    }
                };
            });

            // Welcome page
            app.UseWelcomePage(new WelcomePageOptions{ Path="/wp" });
            */
            app.Run(async (context) =>
            {
                // Uses settings in appsettings.json FIRST, then environment variables, then command line parameters
                // Later overwrites earlier
                // You can also have appsettings.ENVIRONMENTNAME.json which will override the plain one
                // For some reason it isn't reading the global env var, I have to start like this:
                // ASPNETCORE_ENVIRONMENT=Development dotnet run
                var greeting = greeter.GetMessageOfTheDay();
                // await context.Response.WriteAsync($"{greeting}: {env.EnvironmentName}");
                
                // Set the MIME type
                context.Response.ContentType = "text/plain";
                await context.Response.WriteAsync($"Not Found.");
            });
        }
    
        private void ConfigureRoutes(IRouteBuilder routeBuilder)
        {
            // /Home/Index
            // MVC adds "controller" itself
            // Action is the name of the method
            // Supply defaults with =
            // {arg?} means optional
            routeBuilder.MapRoute("Default", "{controller=Home}/{action=Index}/{id?}");
        }

        private IConfiguration _configuration;
    }
}
