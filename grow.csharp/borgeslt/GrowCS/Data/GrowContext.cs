using GrowCS.Models;
using Microsoft.EntityFrameworkCore;

namespace GrowCS.Data
{
    public class GrowContext : DbContext
    {
        public GrowContext(DbContextOptions<GrowContext> options)
            : base(options)
        {
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<GrowData>()
                        .HasKey(e => new
                        {
                            e.Country,
                            e.Indicator,
                            e.Year
                        });

            base.OnModelCreating(modelBuilder);
        }

        public DbSet<GrowData> GrowData { get; set; }
    }
}