<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DataTables;
use Illuminate\Support\Facades\DB;
Use Redirect;
use Illuminate\Support\Facades\Auth;
use Session;

class jsonImportController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        // See getData() function
        if(Auth::check()){
            $path = storage_path()."/challenge.json";
            $items_json = file_get_contents($path);
            $items = json_decode($items_json, true);
            return view('dev.Laraveljson_import_test', compact('items'));
        
        }
  
        return redirect("welcome")->withSuccess('You are not allowed to access');
    }



}
