namespace Pluralsight.Msbuild.Tasks
{
    using Microsoft.Build.Framework;

    public class AddTwoNumbers : ITask
    {
        [Required]
        public double NumberOne { get; set; }
        [Required]
        public double NumberTwo { get; set; }

        [Output]
        public double Result{ get; set; }

        public IBuildEngine BuildEngine { get; set; }
        public ITaskHost HostObject { get; set; }

        public bool Execute()
        {
            Result = NumberOne + NumberTwo;
            BuildEngine.LogMessageEvent(new BuildMessageEventArgs("Added two numbers", "Add", "AddTwoNumbers", MessageImportance.High));

            return true;
        }
    }
}
