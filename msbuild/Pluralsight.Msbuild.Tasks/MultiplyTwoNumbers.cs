namespace Pluralsight.Msbuild.Tasks
{
    using Microsoft.Build.Framework;
    using Microsoft.Build.Utilities;

    public class MultiplyTwoNumbers : Task
    {
        [Required]
        public double NumberOne { get; set; }
        [Required]
        public double NumberTwo { get; set; }

        [Output]
        public double Result { get; set; }

        public override bool Execute()
        {
            Result = NumberOne * NumberTwo;
            Log.LogMessage(MessageImportance.High, "Multiplied two numbers");

            return true;
        }
    }
}
