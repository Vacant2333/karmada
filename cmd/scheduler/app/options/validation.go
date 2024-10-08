/*
Copyright 2021 The Karmada Authors.

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

package options

import (
	"net"
	"strconv"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

// Complete ensures that options are valid and marshals them if necessary.
func (o *Options) Complete() error {
	if len(o.HealthProbeBindAddress) == 0 {
		o.HealthProbeBindAddress = net.JoinHostPort(o.BindAddress, strconv.Itoa(o.SecurePort))
	}
	if len(o.MetricsBindAddress) == 0 {
		o.MetricsBindAddress = net.JoinHostPort(o.BindAddress, strconv.Itoa(o.SecurePort))
	}
	return nil
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() field.ErrorList {
	errs := field.ErrorList{}

	newPath := field.NewPath("Options")
	if o.SchedulerEstimatorPort < 0 || o.SchedulerEstimatorPort > 65535 {
		errs = append(errs, field.Invalid(newPath.Child("SchedulerEstimatorPort"), o.SchedulerEstimatorPort, "must be a valid port between 0 and 65535 inclusive"))
	}
	if o.SchedulerEstimatorTimeout.Duration < 0 {
		errs = append(errs, field.Invalid(newPath.Child("SchedulerEstimatorTimeout"), o.SchedulerEstimatorTimeout, "must be greater than or equal to 0"))
	}

	if o.SchedulerName == "" {
		errs = append(errs, field.Invalid(newPath.Child("SchedulerName"), o.SchedulerName, "should not be empty"))
	}

	return errs
}
