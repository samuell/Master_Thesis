package main

import (
	sp "github.com/scipipe/scipipe"
	spc "github.com/scipipe/scipipe/components"
)

func main() {
	wf := sp.NewWorkflow("DS_CPSign", 4)

	// Create target datasets
	createTargetDatasets := wf.NewProc("create_target_datasets", "python3 ../db_targets.py ../database ../targets_folder; echo 'done' > {o:doneflag}")
	createTargetDatasets.SetOut("doneflag", "log1")

	// Glob the target dataset files into a stream
	targetDirGlobber := spc.NewFileGlobberDependent(wf, "Targets_in_Dir", "./targets_folder/*.json")
	targetDirGlobber.InDependency().From(createTargetDatasets.Out("doneflag"))

	// Balance the target datasets
	balanceTargetDatasets := wf.NewProc("Proc2", "python3 ../balancing_targets.py {i:inpfiles} ../targets_folder; echo 'done' > {o:doneflag}")
	balanceTargetDatasets.In("inpfiles").From(targetDirGlobber.Out())
	balanceTargetDatasets.SetOut("doneflag", "{i:inpfiles|%.json}.done")

	wf.Run()
}
