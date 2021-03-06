/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pkg

import (
	"fmt"
)

import (
	"github.com/apache/dubbo-go/config"
)

type CatService struct {
	GetId   func() (int, error)
	GetName func() (string, error)
	Yell    func() (string, error)
}

func (c *CatService) Reference() string {
	return "CatService"
}

type LionService struct {
	GetId   func() (int, error)
	GetName func() (string, error)
	Yell    func() (string, error)
}

func (l *LionService) Reference() string {
	return "LionService"
}

func init() {
	cat := new(CatService)
	config.SetConsumerService(cat)
	lion := new(LionService)
	config.SetConsumerService(lion)

	config.SetProviderService(&AmericanService{
		cat:  cat,
		lion: lion,
	})
}

type AmericanService struct {
	cat  *CatService
	lion *LionService
}

func (a *AmericanService) Have() (string, error) {
	name, _ := a.cat.GetName()
	return "I'm American and I have a " + name, nil
}

func (a *AmericanService) Hear() (string, error) {
	name, _ := a.lion.GetName()
	yell, _ := a.lion.Yell()
	return fmt.Sprintf("I'm American and I heard a %s yells like %s", name, yell), nil
}

func (a *AmericanService) Reference() string {
	return "AmericanService"
}
