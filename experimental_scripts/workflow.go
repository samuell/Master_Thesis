package main

import (
	sp "github.com/scipipe/scipipe"
	spc "github.com/scipipe/scipipe/components"
)

func main() {
	wf := sp.NewWorkflow("DS_CPSign", 4)

	// Create target datasets
	createTargetDatasets := wf.NewProc("create_target_datasets", "python3 ../db_targets.py ../database {o:targets_dir}; echo 'done' > {o:doneflag}")
	createTargetDatasets.SetOut("targets_dir", "data/target_datasets")
	createTargetDatasets.SetOut("doneflag", "data/target_datasets.done")

	// Glob the target dataset files into a stream
	targetDirGlobber := spc.NewFileGlobberDependent(wf, "glob_target_datasets", "./data/target_datasets/*.json")
	targetDirGlobber.InDependency().From(createTargetDatasets.Out("doneflag"))

	// Balance the target datasets
	balanceTargetDatasets := wf.NewProc("balance_dataset", "python3 ../balancing_targets.py {i:inpfiles} {o:balanced_targets_dir}; echo 'done' > {o:doneflag}")
	balanceTargetDatasets.In("inpfiles").From(targetDirGlobber.Out())
	balanceTargetDatasets.SetOut("balanced_targets_dir", "data/balanced_target_datasets")
	balanceTargetDatasets.SetOut("doneflag", "{i:inpfiles|%.json}.done")

	wf.Run()
}
