"""
Copyright 2021 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

from flask import Blueprint, request
from flask_json import as_json
from magma.protocol_controller.grpc_communication.get_common_rc_response import (
    get_common_bulk_rc_response,
)
from magma.protocol_controller.plugins.cbsd_sas.validators.registration_request import (
    RegistrationRequestSchema,
)

registration_page = Blueprint("registration", __name__)


@registration_page.route('/registration', methods=('POST',))
@as_json
def registration():
    """
    Handle registration route
    """
    return get_common_bulk_rc_response(request, "registrationResponse", RegistrationRequestSchema)
