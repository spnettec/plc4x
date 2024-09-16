/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package task

import (
	"fmt"
	"time"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/debugging"
)

type TaskRequirements interface {
	fmt.Stringer
	ProcessTask() error
	InstallTask(options InstallTaskOptions)
	GetTaskTime() *time.Time
	GetIsScheduled() bool
	SetIsScheduled(isScheduled bool)
}

//go:generate plc4xGenerator -type=Task -prefix=task_
type Task struct {
	*DebugContents   `ignore:"true"`
	TaskRequirements `ignore:"true"`
	taskTime         *time.Time
	isScheduled      bool
}

func NewTask(taskRequirements TaskRequirements, opts ...func(*Task)) *Task {
	t := &Task{TaskRequirements: taskRequirements}
	t.DebugContents = NewDebugContents(t, "taskTime", "isScheduled")
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func (t *Task) GetDebugAttr(attr string) any {
	switch attr {
	case "taskTime":
		if t.taskTime != nil {
			return *t.taskTime
		}
	case "isScheduled":
		return t.isScheduled
	}
	return nil
}

func (t *Task) InstallTask(options InstallTaskOptions) {
	when := options.When
	delta := options.Delta
	// check for delta from now
	if when == nil && delta != nil {
		_when := _taskManager.GetTime().Add(*delta)
		when = &_when
	}

	// fallback to the initial value
	if when == nil {
		_when := t.taskTime
		when = _when
	}
	if when == nil {
		panic("schedule missing, use zero for 'now'")
	}
	t.taskTime = when

	// pass along to the task manager
	_taskManager.InstallTask(t.TaskRequirements)
}

func (t *Task) SuspendTask() {
	_taskManager.SuspendTask(t.TaskRequirements)
}

func (t *Task) Resume() {
	_taskManager.ResumeTask(t.TaskRequirements)
}

func (t *Task) GetTaskTime() *time.Time {
	return t.taskTime
}

func (t *Task) GetIsScheduled() bool {
	return t.isScheduled
}

func (t *Task) SetIsScheduled(isScheduled bool) {
	t.isScheduled = isScheduled
}
