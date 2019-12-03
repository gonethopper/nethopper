// MIT License

// Copyright (c) 2019 gonethopper

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// * @Author: ankye
// * @Date: 2019-06-24 11:07:19
// * @Last Modified by:   ankye
// * @Last Modified time: 2019-06-24 11:07:19

package logic

import (
	"time"

	"github.com/gonethopper/nethopper/examples/simple_server/common"
	"github.com/gonethopper/nethopper/server"
)

// LogicService struct to define service
type LogicService struct {
	server.BaseContext
}

// LogicServiceCreate  service create function
func LogicServiceCreate() (server.Service, error) {
	return &LogicService{}, nil
}

// UserData service custom option, can you store you data and you must keep goruntine safe
func (s *LogicService) UserData() int32 {
	return 0
}

// Setup init custom service and pass config map to service
// config
// m := map[string]interface{}{
//  "queueSize":1000,
// }
func (s *LogicService) Setup(m map[string]interface{}) (server.Service, error) {
	s.RegisterHandler(common.CallIDLoginCmd, LoginHandler)
	return s, nil
}

//Reload reload config
func (s *LogicService) Reload(m map[string]interface{}) error {
	return nil
}

// OnRun goruntine run and call OnRun , always use ServiceRun to call this function
func (s *LogicService) OnRun(dt time.Duration) {
	for i := 0; i < 128; i++ {
		m, err := s.MQ().AsyncPop()
		if err != nil {
			break
		}
		obj := m.(*server.CallObject)

		go server.Processor(s, obj)

	}

}

// Stop goruntine
func (s *LogicService) Stop() error {
	return nil
}

// Call async send message to service
func (s *LogicService) Call(option int32, obj *server.CallObject) error {
	if err := s.MQ().AsyncPush(obj); err != nil {
		server.Error(err.Error())
	}
	return nil
}

// PushBytes async send string or bytes to queue
func (s *LogicService) PushBytes(option int32, buf []byte) error {
	return nil
}
