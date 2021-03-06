package wasmtime

import "testing"

func TestModule(t *testing.T) {
	_, err := NewModule(NewStore(NewEngine()), []byte{})
	if err == nil {
		panic("expected an error")
	}
	_, err = NewModule(NewStore(NewEngine()), []byte{1})
	if err == nil {
		panic("expected an error")
	}
}

func TestModuleValidate(t *testing.T) {
	if ModuleValidate(NewStore(NewEngine()), []byte{}) == nil {
		panic("expected an error")
	}
	if ModuleValidate(NewStore(NewEngine()), []byte{1}) == nil {
		panic("expected an error")
	}
	wasm, err := Wat2Wasm(`(module)`)
	if err != nil {
		panic(err)
	}
	if ModuleValidate(NewStore(NewEngine()), wasm) != nil {
		panic("expected valid module")
	}
}

func TestModuleImports(t *testing.T) {
	wasm, err := Wat2Wasm(`
          (module
            (import "" "f" (func))
            (import "a" "g" (global i32))
            (import "" "" (table 1 funcref))
            (import "" "" (memory 1))
          )
        `)
	if err != nil {
		panic(err)
	}
	module, err := NewModule(NewStore(NewEngine()), wasm)
	if err != nil {
		panic(err)
	}
	imports := module.Imports()
	if len(imports) != 4 {
		panic("wrong number of imports")
	}
	if imports[0].Module() != "" {
		panic("wrong import module")
	}
	if imports[0].Name() != "f" {
		panic("wrong import name")
	}
	if imports[0].Type().FuncType() == nil {
		panic("wrong import type")
	}
	if len(imports[0].Type().FuncType().Params()) != 0 {
		panic("wrong import type")
	}
	if len(imports[0].Type().FuncType().Results()) != 0 {
		panic("wrong import type")
	}

	if imports[1].Module() != "a" {
		panic("wrong import module")
	}
	if imports[1].Name() != "g" {
		panic("wrong import name")
	}
	if imports[1].Type().GlobalType() == nil {
		panic("wrong import type")
	}
	if imports[1].Type().GlobalType().Content().Kind() != KindI32 {
		panic("wrong import type")
	}

	if imports[2].Module() != "" {
		panic("wrong import module")
	}
	if imports[2].Name() != "" {
		panic("wrong import name")
	}
	if imports[2].Type().TableType() == nil {
		panic("wrong import type")
	}
	if imports[2].Type().TableType().Element().Kind() != KindFuncref {
		panic("wrong import type")
	}

	if imports[3].Module() != "" {
		panic("wrong import module")
	}
	if imports[3].Name() != "" {
		panic("wrong import name")
	}
	if imports[3].Type().MemoryType() == nil {
		panic("wrong import type")
	}
	if imports[3].Type().MemoryType().Limits().Min != 1 {
		panic("wrong import type")
	}
}

func TestModuleExports(t *testing.T) {
	wasm, err := Wat2Wasm(`
          (module
            (func (export "f"))
            (global (export "g") i32 (i32.const 0))
            (table (export "t") 1 funcref)
            (memory (export "m") 1)
          )
        `)
	if err != nil {
		panic(err)
	}
	module, err := NewModule(NewStore(NewEngine()), wasm)
	if err != nil {
		panic(err)
	}
	exports := module.Exports()
	if len(exports) != 4 {
		panic("wrong number of exports")
	}
	if exports[0].Name() != "f" {
		panic("wrong export name")
	}
	if exports[0].Type().FuncType() == nil {
		panic("wrong export type")
	}
	if len(exports[0].Type().FuncType().Params()) != 0 {
		panic("wrong export type")
	}
	if len(exports[0].Type().FuncType().Results()) != 0 {
		panic("wrong export type")
	}

	if exports[1].Name() != "g" {
		panic("wrong export name")
	}
	if exports[1].Type().GlobalType() == nil {
		panic("wrong export type")
	}
	if exports[1].Type().GlobalType().Content().Kind() != KindI32 {
		panic("wrong export type")
	}

	if exports[2].Name() != "t" {
		panic("wrong export name")
	}
	if exports[2].Type().TableType() == nil {
		panic("wrong export type")
	}
	if exports[2].Type().TableType().Element().Kind() != KindFuncref {
		panic("wrong export type")
	}

	if exports[3].Name() != "m" {
		panic("wrong export name")
	}
	if exports[3].Type().MemoryType() == nil {
		panic("wrong export type")
	}
	if exports[3].Type().MemoryType().Limits().Min != 1 {
		panic("wrong export type")
	}
}
