/*
 * Copyright 2021 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package chans

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/Comcast/plax/dsl"
)

func TestSQS(t *testing.T) {

	// Can use (say) https://github.com/p4tin/goaws to run a
	// local/mock SQS.  That goaws apparently needs a config file
	// with an account id in order to return valid XML in some
	// cases.  The file 'goaws.config' does that.  Example:
	//
	//  goaws -config goaws.config
	//
	// If this test fails to talk to an SQS, the test is skipped.
	//
	//
	// Use the AWS CLI to create a queue for testing:
	//
	// aws --endpoint-url http://localhost:4100 sqs create-queue --queue-name plaxtest

	endpoint := "http://localhost:4100"

	opts := SQSOpts{
		Endpoint:        endpoint,
		QueueURL:        endpoint + "/123456789/plaxtest",
		MsgDelaySeconds: true,
	}

	ctx := dsl.NewCtx(context.Background())

	c, err := NewSQSChan(ctx, opts)
	if err != nil {
		t.Fatal(err)
	}

	skip := func(err error) {
		t.Skipf("skipping SQS test (%s)", err)
	}

	if err = c.Open(ctx); err != nil {
		skip(err)
	}

	defer c.Close(ctx)

	then := time.Now().UTC().Format(time.RFC3339Nano)

	{
		m := dsl.Msg{
			Payload: map[string]interface{}{
				"t":            then,
				"want":         "tacos",
				"DelaySeconds": 2,
			},
		}

		if err = c.Pub(ctx, m); err != nil {
			skip(err)
		}
	}

	{
		ch := c.Recv(ctx)
		msg := <-ch
		js, err := json.MarshalIndent(&msg, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("recv\n%s\n", js)

		if m, is := msg.Payload.(map[string]interface{}); is {
			then0 := m["t"]
			if then0 != then {
				t.Fatalf("%v != %v", then0, then)
			}
		} else {
			t.Fatalf("%T", msg.Payload)
		}
	}

}
