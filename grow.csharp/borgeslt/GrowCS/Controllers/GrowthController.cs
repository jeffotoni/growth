using System;
using System.Text;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using GrowCS.Data;
using GrowCS.Models;

namespace GrowCS.Controllers
{
    [ApiController]
    [Route("api/v1/growth")]
    [Produces("application/json")]
    public class GrowController : ControllerBase
    {
        private readonly GrowContext _context;
        public GrowController(GrowContext context)
        {
            _context = context;
        }

        [HttpGet]
        [Route("/ping")]
        public ActionResult Ping()
        {
            return Ok(new { msg = "pong❤" });
        }

        [HttpPost]
        [Produces("application/json")]
        public async Task<ActionResult> Post([FromBody] Object item)
        {   
            if (item == null) return BadRequest();
            Console.WriteLine(":...debug in Json...:");

            var result = await _context.GrowData.CountAsync();
            if (result > 0) {
                return Ok(new { msg = "In progress" });
            }

            string sjson = item.ToString();
            var data = JsonSerializer.Deserialize<GrowData[]>(sjson);
            _context.GrowData.AddRange(data);
            await _context.SaveChangesAsync();
            return Ok(new { msg = "In progress" });
        }

        [HttpPost]
        [Route("/form")]
        public async Task<ActionResult> PostForm([FromForm] JsonFile jsonFile)
        {
            var file = jsonFile.File;
            if (file == null || file.Length == 0)
            {

                return BadRequest(new { msg = "error in your json" });
            }

            using var reader = new StreamReader(file.OpenReadStream());
            var content = await reader.ReadToEndAsync();
            var data = JsonSerializer.Deserialize<GrowData[]>(content);

            _context.GrowData.AddRange(data);
            await _context.SaveChangesAsync();

            return Ok(new { msg = "In progress" });
        }

        [HttpGet]
        [Route("post/status")]
        [Produces("application/json")]
        public async Task<ActionResult> Status()
        {
            var result = await _context.GrowData.FirstOrDefaultAsync(d => d.Country == "BRZ" && d.Indicator == "NGDP_R" && d.Year == 2002);
            if (result == null)
            {
                return BadRequest(new { msg = "not finished" });
            }

            return Ok(new
            {
                msg = "Complete",
                test_value = $"{result.Value:f2}",
                count = _context.GrowData.Count()
            });
        }

        [HttpGet]
        [Route("{country}/{indicator}/{year}")]
        [Produces("application/json")]
        public async Task<ActionResult> Get(string country, string indicator, int year)
        {
           Console.WriteLine(":...debug in Json...:" + country);
            var result = await _context.GrowData.FirstOrDefaultAsync(d => d.Country == country.ToUpper()
            && d.Indicator == indicator.ToUpper() && d.Year == year);
            if (result == null)
            {
                return BadRequest(new { msg = "error in path url" });
            }

            return Ok(new
            {
                result.Country,
                result.Indicator,
                result.Year,
                Value = $"{result.Value:f2}",
            });
        }

        [HttpGet]
        [Route("size")]
        [Produces("application/json")]
        public async Task<ActionResult> GetSize()
        {
            var result = await _context.GrowData.CountAsync();
            return Ok(new { size = result });
        }

        [HttpPut]
        [Route("{country}/{indicator}/{year}")]
        [Produces("application/json")]        
        public async Task<ActionResult> Put(string country, string indicator, int year, [FromBody] Object item)
        {   
            var action = "Updated";
            var result = await _context.GrowData.FirstOrDefaultAsync(d => d.Country == country.ToUpper() 
            && d.Indicator == indicator.ToUpper() && d.Year == year);

            string sjson = item.ToString();
            var data = JsonSerializer.Deserialize<GrowData>(sjson);

            if (result == null)
            {
                if (item == null) return BadRequest();
               
                Console.WriteLine(":...debug in item Json...:"+data.Value);

                _context.GrowData.Add(new GrowData
                {
                    Country = country.ToUpper(),
                    Indicator = indicator.ToUpper(),
                    Year = year,
                    Value = data.Value
                });
                action = "Inserted";
            }
            else
            {
                result.Value = data.Value;
                _context.Update(result);
            }

            await _context.SaveChangesAsync();

            return Ok(new
            {
                msg = action
            });
        }

        [HttpDelete]
        [Route("{country}/{indicator}/{year}")]
        [Produces("application/json")]
        public async Task<ActionResult> Delete(string country, string indicator, int year)
        {
            var result = await _context.GrowData.FirstOrDefaultAsync(d => d.Country == country.ToUpper() && d.Indicator == indicator.ToUpper() && d.Year == year);
            if (result == null)
            {
                return BadRequest(new { msg = "error in path url" });
            }

            _context.GrowData.Remove(result);
            await _context.SaveChangesAsync();
            return Ok(new { msg = "Deleted" });
        }
    }
}