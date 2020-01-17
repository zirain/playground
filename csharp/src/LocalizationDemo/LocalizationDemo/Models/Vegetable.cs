using System.ComponentModel.DataAnnotations;

namespace LocalizationDemo.Models
{
    public enum Vegetable
    {
        [Display(Name = "Kale")]
        Kale,

        [Display(Name = "Spinach")]
        Spinach,

        [Display(Name = "Cauliflower")]
        Cauliflower
    }
}
