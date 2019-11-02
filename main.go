/*
Copyright © 2019 ichbinfrog

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (

	// "github.com/ichbinfrog/petname/cmd"

	"fmt"
	"os"

	"github.com/ichbinfrog/petname/pkg/dict"
)

func main() {
	// cmd.Execute()
	tree := &dict.Tree{}
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(20)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(19)

	dict.Print(os.Stdout, tree.Root, 2, '└')
	fmt.Println(tree.Search([]int{1, 5, 8, 20, 19, 30}))
}
