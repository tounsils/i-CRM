<?php
// App\Models\user.php
namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class user extends Model
{
    protected $table = 'appusers';
    use HasFactory;
    public $fillable = [
        'first_name',
        'last_name',
        'email',     
        'phone',
    ];    
}
