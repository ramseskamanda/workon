package system

type InternalIndex struct {
}

func (index *InternalIndex) FindProjectByName(project string) (string, error) {
	// files, err := ioutil.ReadDir(DIRECTORY)
	// if err != nil {
	// 	return []fs.FileInfo{}, nil
	// }

	// filtered := []fs.FileInfo{}
	// for _, f := range files {
	// 	if f.IsDir() && strings.Contains(f.Name(), project) {
	// 		filtered = append(filtered, f)
	// 	}
	// }
	// return filtered, nil

	return "", nil
}
