extern crate js_snap;


#[repr(C)]
pub struct Instance {

}

#[no_mangle]
pub extern fn client_snap_init() {
  js_snap::js_snap_init();
}

#[no_mangle]
pub extern fn client_snap_instance_new<'a>(
    export_name: *const std::os::raw::c_char) -> *mut js_snap::Instance<'a>
{
  let bytes: &'static [u8] = include_bytes!(concat!(env!("OUT_DIR"), "/ssr.bin"));
  js_snap::js_snap_instance_from_snapshot(
    bytes.as_ptr(),
    bytes.len(),
    export_name)
}

#[no_mangle]
pub extern fn client_snap_instance_delete<'a>(instance: *mut js_snap::Instance<'a>) {
  js_snap::js_snap_instance_delete(instance)
}

#[no_mangle]
pub extern fn client_snap_instance_call<'a>(
  instance: *mut js_snap::Instance<'a>,
  name: *const std::os::raw::c_char,
  params: *const std::os::raw::c_char,
  result_ptr: *mut *const std::os::raw::c_char,
  result_len: *mut i32,
) {
  js_snap::js_snap_instance_call(instance, name, params, result_ptr, result_len)
}