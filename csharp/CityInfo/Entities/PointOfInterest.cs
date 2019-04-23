using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace CityInfo.Entities
{
    public class PointOfInterest
    {
        // A field with name Id is automatically the primary key (PointOfInterestId would work too)
        // But we should annotate anyway
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id { get; set; }
        // Apply constraints at lowest possible level (at the database)
        [Required]
        [MaxLength(50)]
        public string Name { get; set; }
        
        // This will be auto-discovered as a navigation property
        // But we annotate anyway
        [ForeignKey("CityId")]
        public City City { get; set; }
        // Auto-discovered as foreign key (but we annotated anyway)
        public int CityId { get; set; }
    }
}
