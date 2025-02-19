package v1alpha1

import (
	"github.com/weaveworks/ignite/pkg/apis/ignite"
	meta "github.com/weaveworks/ignite/pkg/apis/meta/v1alpha1"
	"k8s.io/apimachinery/pkg/conversion"
)

// Convert_ignite_VMSpec_To_v1alpha1_VMSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_ignite_VMSpec_To_v1alpha1_VMSpec(in *ignite.VMSpec, out *VMSpec, s conversion.Scope) error {
	// VMSpecStorage are not supported by v1alpha1, so just ignore the warning by calling this manually
	return autoConvert_ignite_VMSpec_To_v1alpha1_VMSpec(in, out, s)
}

// Convert_ignite_VMSpec_To_v1alpha1_VMSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_v1alpha1_VMSpec_To_ignite_VMSpec(in *VMSpec, out *ignite.VMSpec, s conversion.Scope) error {
	// VMSpecStorage is not supported by v1alpha1, so just ignore the warning by calling this manually
	return autoConvert_v1alpha1_VMSpec_To_ignite_VMSpec(in, out, s)
}

// Convert_ignite_VMStatus_To_v1alpha1_VMStatus calls the autogenerated conversion function along with custom conversion logic
func Convert_ignite_VMStatus_To_v1alpha1_VMStatus(in *ignite.VMStatus, out *VMStatus, s conversion.Scope) error {
	if err := autoConvert_ignite_VMStatus_To_v1alpha1_VMStatus(in, out, s); err != nil {
		return err
	}

	// Convert in.Running to out.State
	if in.Running {
		out.State = VMStateRunning
	} else {
		out.State = VMStateStopped
	}

	return nil
}

// Convert_v1alpha1_VMStatus_To_ignite_VMStatus calls the autogenerated conversion function along with custom conversion logic
func Convert_v1alpha1_VMStatus_To_ignite_VMStatus(in *VMStatus, out *ignite.VMStatus, s conversion.Scope) error {
	if err := autoConvert_v1alpha1_VMStatus_To_ignite_VMStatus(in, out, s); err != nil {
		return err
	}

	// Convert in.State to out.Running
	out.Running = in.State == VMStateRunning

	return nil
}

// Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource calls the autogenerated conversion function along with custom conversion logic
func Convert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(in *ignite.OCIImageSource, out *OCIImageSource, s conversion.Scope) error {
	if err := autoConvert_ignite_OCIImageSource_To_v1alpha1_OCIImageSource(in, out, s); err != nil {
		return err
	}

	// If the OCI content ID is local, i.e. not available from a repository,
	// populate the ID field of v1alpha1.OCIImageSource. Otherwise add the
	// the repo digest of the ID as the only digest for v1alpha1.
	if in.ID.Local() {
		out.ID = in.ID.Digest().String()
	} else {
		out.RepoDigests = []string{in.ID.RepoDigest().String()}
	}

	return nil
}

// Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource calls the autogenerated conversion function along with custom conversion logic
func Convert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(in *OCIImageSource, out *ignite.OCIImageSource, s conversion.Scope) (err error) {
	if err = autoConvert_v1alpha1_OCIImageSource_To_ignite_OCIImageSource(in, out, s); err != nil {
		return err
	}

	// By default parse the OCI content ID from the Docker image ID
	contentRef := in.ID
	if len(in.RepoDigests) > 0 {
		// If the image has Repo digests, use the first one of them to parse
		// the fully qualified OCI image name and digest. All of the digests
		// point to the same contents, so it doesn't matter which one we use
		// here. It will be translated to the right content by the runtime.
		contentRef = in.RepoDigests[0]
	}

	// Parse the OCI content ID based on the available reference
	out.ID, err = meta.ParseOCIContentID(contentRef)
	return
}

func Convert_v1alpha1_OCIClaim_To_ignite_OCI(in *OCIImageClaim, out *meta.OCIImageRef) error {
	// Convert the old OCIImageClaim format to meta.OCIImageRef
	// by extracting in.Ref and ignoring in.Type
	*out = in.Ref

	return nil
}

func Convert_ignite_OCI_To_v1alpha1_OCIClaim(in *meta.OCIImageRef, out *OCIImageClaim) error {
	// Convert meta.OCIImageRef to the old OCIImageClaim format,
	// set out.Ref to the OCIImageRef and out.Type to the default
	// ImageSourceTypeDocker
	out.Ref = *in
	out.Type = ImageSourceTypeDocker

	return nil
}

// Convert_ignite_ImageSpec_To_v1alpha1_ImageSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_ignite_ImageSpec_To_v1alpha1_ImageSpec(in *ignite.ImageSpec, out *ImageSpec, s conversion.Scope) error {
	if err := autoConvert_ignite_ImageSpec_To_v1alpha1_ImageSpec(in, out, s); err != nil {
		return err
	}

	return Convert_ignite_OCI_To_v1alpha1_OCIClaim(&in.OCI, &out.OCIClaim)
}

// Convert_v1alpha1_ImageSpec_To_ignite_ImageSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_v1alpha1_ImageSpec_To_ignite_ImageSpec(in *ImageSpec, out *ignite.ImageSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha1_ImageSpec_To_ignite_ImageSpec(in, out, s); err != nil {
		return err
	}

	return Convert_v1alpha1_OCIClaim_To_ignite_OCI(&in.OCIClaim, &out.OCI)
}

// Convert_ignite_KernelSpec_To_v1alpha1_KernelSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_ignite_KernelSpec_To_v1alpha1_KernelSpec(in *ignite.KernelSpec, out *KernelSpec, s conversion.Scope) error {
	if err := autoConvert_ignite_KernelSpec_To_v1alpha1_KernelSpec(in, out, s); err != nil {
		return err
	}

	return Convert_ignite_OCI_To_v1alpha1_OCIClaim(&in.OCI, &out.OCIClaim)
}

// Convert_v1alpha1_KernelSpec_To_ignite_KernelSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_v1alpha1_KernelSpec_To_ignite_KernelSpec(in *KernelSpec, out *ignite.KernelSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha1_KernelSpec_To_ignite_KernelSpec(in, out, s); err != nil {
		return err
	}

	return Convert_v1alpha1_OCIClaim_To_ignite_OCI(&in.OCIClaim, &out.OCI)
}

// Convert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec(in *ignite.VMImageSpec, out *VMImageSpec, s conversion.Scope) error {
	if err := autoConvert_ignite_VMImageSpec_To_v1alpha1_VMImageSpec(in, out, s); err != nil {
		return err
	}

	return Convert_ignite_OCI_To_v1alpha1_OCIClaim(&in.OCI, &out.OCIClaim)
}

// Convert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec(in *VMImageSpec, out *ignite.VMImageSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha1_VMImageSpec_To_ignite_VMImageSpec(in, out, s); err != nil {
		return err
	}

	return Convert_v1alpha1_OCIClaim_To_ignite_OCI(&in.OCIClaim, &out.OCI)
}

// Convert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec(in *ignite.VMKernelSpec, out *VMKernelSpec, s conversion.Scope) error {
	if err := autoConvert_ignite_VMKernelSpec_To_v1alpha1_VMKernelSpec(in, out, s); err != nil {
		return err
	}

	return Convert_ignite_OCI_To_v1alpha1_OCIClaim(&in.OCI, &out.OCIClaim)
}

// Convert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec(in *VMKernelSpec, out *ignite.VMKernelSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha1_VMKernelSpec_To_ignite_VMKernelSpec(in, out, s); err != nil {
		return err
	}

	return Convert_v1alpha1_OCIClaim_To_ignite_OCI(&in.OCIClaim, &out.OCI)
}

// Convert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec calls the autogenerated conversion function and custom conversion logic
func Convert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec(in *VMNetworkSpec, out *ignite.VMNetworkSpec, s conversion.Scope) error {
	// .Spec.Network.Mode is not tracked in the v1alpha2 and newer, so there's no extra conversion logic
	return autoConvert_v1alpha1_VMNetworkSpec_To_ignite_VMNetworkSpec(in, out, s)
}
