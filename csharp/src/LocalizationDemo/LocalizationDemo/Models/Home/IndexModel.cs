using System.ComponentModel.DataAnnotations;

namespace LocalizationDemo.Models.Home
{
    public class IndexModel
    {
        [Display(Name = "Title.Display")]
        public string Title { get; set; }

        public Vegetable Vegetable { get; set; } = Vegetable.Cauliflower;
    }
}
